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
import ErrorMessage from "../../../koksmat/components/errormessage";

function Item(props: { item: Item }) {
  const { item } = props;
  return (
    <div>
      <a href={"/apps/application/" + item.id} className="text-blue-600">
        {item.displayname}
      </a>
    </div>
  );
}

export default function Owners() {
  const { data, error, isLoading } = useService<Output>(
    "magic-devices.application",
    ["search", "%voice%"],
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

  const owners: Output = data;
  return (
    <div>
      <div className="space-y-4">
        <h1 className="text-2xl font-bold  sm:text-4xl md:text-5xl">
          Applications
        </h1>
      </div>

      {owners.items
        .sort((a, b) => a.displayname.localeCompare(b.displayname))
        .map((item, index) => (
          <div key={index}>
            <Item item={item} />
          </div>
        ))}
    </div>
  );
}
