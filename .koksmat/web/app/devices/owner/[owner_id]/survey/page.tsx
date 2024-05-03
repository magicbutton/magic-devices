"use client";

export interface Root {
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
  survey_responses: SurveyResponse[];
}

export interface SurveyResponse {
  id: number;
  created_at: string;
  updated_at: string;
  tenant: string;
  name: string;
  description: string;
  url: string;
  key: string;
  displayname: string;
  respondent_id: number;
  survey_id: number;
  application_id: number;
  questions: any;
  answers: any;
  answer1: string;
  answer2: string;
  answer3: string;
  answer4: string;
  answer5: string;
  answer6: string;
  answer7: string;
  answer8: string;
  answer9: string;
  truefalse1: boolean;
  truefalse2: boolean;
  truefalse3: boolean;
  datetime1: string;
  datetime2: string;
  datetime3: string;
  number1: number;
  number2: number;
  number3: number;
}

import { useService } from "@/koksmat/useservice";
import ErrorMessage from "@/koksmat/components/errormessage";
import { useEffect } from "react";
import Link from "next/link";

function Survey(props: { response: SurveyResponse }) {
  const { respondent_id, survey_id } = props.response;
  const { data, error, isLoading } = useService<Root>(
    "magic-devices.survey",
    ["read", survey_id.toString()],
    "",
    600,
    "x"
  );
  if (error) {
    return <ErrorMessage message={error} />;
  }
  const survey: Root = data;
  if (survey) {
    console.log(survey);
  }
  return (
    <Link
      className="hover:underline cursor-pointer text-blue-500 dark:text-blue-400 text-xl"
      href={`/apps/owner/${respondent_id}/survey/${survey_id}`}
    >
      <div>{survey?.displayname}</div>
    </Link>
  );
}

export default function ResponseSurveys(props: {
  params: { owner_id: string };
}) {
  const { owner_id } = props.params;
  const { data, error, isLoading } = useService<Root>(
    "magic-devices.person",
    ["withsurveys", owner_id],
    "",
    600,
    "x"
  );
  if (error) {
    return <ErrorMessage message={error} />;
  }

  const responses: Root = data;
  return (
    <div>
      <div className="space-y-4">
        <h1 className="text-2xl font-bold  sm:text-4xl md:text-5xl">Surveys</h1>
        <p className="text-gray-500 dark:text-gray-400 max-w-[650px] text-lg md:text-xl">
          Here is a list of the surveys which you are involved in
        </p>
      </div>
      <div className="mt-6">
        {!data && <div>Loading...</div>}

        {responses?.survey_responses
          .sort((a, b) => a.displayname.localeCompare(b.displayname))
          .map((response, index) => (
            <div key={index}>
              <Survey response={response} />
            </div>
          ))}
      </div>
    </div>
  );
}
