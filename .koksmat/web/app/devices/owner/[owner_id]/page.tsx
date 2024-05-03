"use client";

export interface Root {
  hasError: boolean;
  errorMessage: string;
  data: Data;
}

export interface Data {
  id: number;
  created_at: string;
  updated_at: string;
  tenant: string;
  name: string;
  description: string;
  unique_person_id: string;
  displayname: string;
  email: string;
  applications: Application[];
}

export interface Application {
  id: number;
  created_at: string;
  updated_at: string;
  tenant: string;
  name: string;
  description: string;
  key: string;
  displayname: string;
  imported_ownername: string;
  owner: Owner;
}

export interface Owner {
  id: string;
  name: string;
  entity: string;
}

import { useService } from "@/koksmat/useservice";
import ErrorMessage from "../../../../koksmat/components/errormessage";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";

import Link from "next/link";
import { Button } from "@/components/ui/button";

function App(props: { app: Application }) {
  return (
    <div>
      <div>{props.app.displayname} </div>
    </div>
  );
}

export default function OwnerPage(props: { params: { owner_id: string } }) {
  const { owner_id } = props.params;
  const { data, error, isLoading } = useService<Data>(
    "magic-devices.person",
    ["read", owner_id],
    "",
    600,
    "x"
  );

  if (error) {
    return <ErrorMessage message={error} />;
  }
  if (!data) {
    return <div>Loading...</div>;
  }

  const owner: Data = data;
  return (
    <div>
      <div className="space-y-4">
        <h1 className="text-2xl font-bold  sm:text-4xl md:text-5xl">
          Applications owned by {owner?.displayname}
        </h1>
        <p className="text-gray-500 dark:text-gray-400 max-w-[650px] text-lg md:text-xl">
          Select the applications you own from the list below. This will help us
          verify your account and provide the best experience.
        </p>
      </div>
      <div className="mt-4 flex">
        <div>
          {owner?.applications
            .sort((a, b) => a.displayname.localeCompare(b.displayname))
            .map((app, index) => (
              <div key={index}>
                <App app={app} />
              </div>
            ))}
        </div>
        <div className="grow"></div>
        <div>
          <Link href={`/apps/owner/${owner_id}/survey`}>
            <Button variant="default">Surveys</Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
