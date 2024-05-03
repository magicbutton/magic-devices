"use client";

export interface Output {
  totalPages: number;
  totalItems: number;
  currentPage: number;
  items: Item[];
}

export interface Item {
  id: number;
  created_at: string;
  updated_at: string;
  tenant: string;
  name: string;
  description: string;
  unique_person_id: string;
  displayname: string;
  email: string;
  applications: any;
}

import { useService } from "@/koksmat/useservice";
import ErrorMessage from "@/koksmat/components/errormessage";

function Owner(props: { owner: Item }) {
  const { owner } = props;
  return (
    <div>
      <a href={"/apps/owner/" + owner.id} className="text-blue-600">
        {owner.displayname}
      </a>
    </div>
  );
}

export default function Owners() {
  const { data, error, isLoading } = useService<Output>(
    "magic-devices.person",
    ["search", "%"],
    "",
    600,
    "x"
  );
  if (error) {
    return <ErrorMessage message={error} />;
  }

  const owners: Output = data;
  return (
    <div>
      <div className="space-y-4">
        <h1 className="text-2xl font-bold  sm:text-4xl md:text-5xl">
          Applications owners
        </h1>
        <p className="text-gray-500 dark:text-gray-400 max-w-[650px] text-lg md:text-xl">
          Select the applications you own from the list below. This will help us
          verify your account and provide the best experience.
        </p>
      </div>
      {!data && <div>Loading...</div>}

      {owners?.items
        .sort((a, b) => a.displayname.localeCompare(b.displayname))
        .map((owner, index) => (
          <div key={index}>
            <Owner owner={owner} />
          </div>
        ))}
    </div>
  );
}
