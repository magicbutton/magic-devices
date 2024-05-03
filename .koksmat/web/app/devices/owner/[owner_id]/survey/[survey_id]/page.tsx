"use client";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export interface Survey {
  id: number;
  created_at: string;
  updated_at: string;
  tenant: string;
  name: string;
  description: string;
  url: string;
  key: string;
  displayname: string;
  owner_id: number;
  campaign_id: number;
  question1: string;
  question2: string;
  question3: string;
  question4: string;
  question5: string;
  question6: string;
  question7: string;
  question8: string;
  question9: string;
  truefalse1: string;
  truefalse2: string;
  truefalse3: string;
  datetime1: string;
  datetime2: string;
  datetime3: string;
  number1: string;
  number2: string;
  number3: string;
  questions: any;
}

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
  responsedate: string;
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
import { Fragment, useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { run } from "@/koksmat/magicservices";

function YesNo(props: {
  id: string;
  question: string;
  checked: boolean;
  onChange: (value: boolean) => void;
}) {
  const { id, question, checked, onChange } = props;
  return (
    <div className="space-y-3">
      <div>{question}</div>
      <div className="flex items-center space-x-2">
        <Switch
          className="pointer "
          id={"survey_response" + id}
          checked={checked}
          onCheckedChange={onChange}
        />
        <Label className="pointer " htmlFor={"survey_response" + id}>
          Turn on to answer yes or off to answer no.
        </Label>
      </div>
    </div>
  );
}

export default function ResponseSurveys(props: {
  params: { owner_id: string; survey_id: string };
}) {
  const { owner_id, survey_id } = props.params;
  const { data, error, isLoading } = useService<Root>(
    "magic-devices.person",
    ["withsurveyresponses", owner_id, survey_id],
    "",
    600,
    "x"
  );
  const surveyItem = useService<Survey>(
    "magic-devices.survey",
    ["read", survey_id],
    "",
    600,
    "x"
  );

  const [version, setversion] = useState(0);
  const survey: Survey = surveyItem.data;
  const [allanswered, setallanswered] = useState(false);

  const [responses, setresponses] = useState<SurveyResponse[]>([]);
  const submit = async () => {
    const sendResult = await run(
      "magic-devices.surveyresponse",
      ["submit", JSON.stringify(responses)],
      "",
      60,
      "x"
    );

    if (sendResult.hasError) {
      console.log(sendResult.errorMessage);
    } else {
      console.log("success");
    }
  };

  useEffect(() => {
    if (!survey) return;

    setresponses(data?.survey_responses);
  }, [survey]);
  if (error) {
    return <ErrorMessage message={error} />;
  }

  const person = data;
  return (
    <div>
      <div className="space-y-4">
        <h1 className="text-2xl font-bold  sm:text-4xl md:text-5xl">
          {survey?.displayname}
        </h1>
        <p className="text-gray-500 dark:text-gray-400 max-w-[650px] text-lg md:text-xl">
          Here is a list of apps recorded as owned by{" "}
          <b> {person?.displayname} </b>
        </p>
      </div>
      <div className="flex items-center space-x-2 mt-6">
        <Button
          variant="secondary"
          onClick={() => {
            responses.map((response) => {
              response.truefalse1 = true;
            });

            setversion(version + 1);
            setresponses(responses);
          }}
        >
          Yes to all
        </Button>
        <Button
          variant="secondary"
          onClick={() => {
            responses.map((response) => {
              response.truefalse1 = false;
            });

            setversion(version + 1);
            setresponses(responses);
          }}
        >
          No to all
        </Button>
        <div className="grow"></div>
        <div className="flex items-center space-x-2">
          <Switch
            id="airplane-mode"
            checked={allanswered}
            onCheckedChange={() => {
              setallanswered(!allanswered);
            }}
          />
          <Label htmlFor="airplane-mode">All have been answered</Label>
        </div>
        <Button
          disabled={!allanswered}
          onClick={() => {
            submit();
          }}
        >
          Send
        </Button>
      </div>
      <div className="mt-5 border rounded-lg p-10">
        {!data && <div>Loading...</div>}

        <div>
          <div className="hidden">{version}</div>
          <div className="columns-1 md:columns-2 xl:columns-3 space-y-3">
            {responses
              .sort((a, b) => a.displayname.localeCompare(b.displayname))
              .map((response, index) => (
                <Card key={index}>
                  <CardHeader>
                    <CardTitle>{response.displayname}</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <YesNo
                      id={response.id.toString()}
                      question={survey.question1}
                      checked={response.truefalse1}
                      onChange={(value) => {
                        response.truefalse1 = value;
                        setversion(version + 1);
                        setresponses(responses);
                      }}
                    />
                    <YesNo
                      id={"q2" + response.id.toString()}
                      question={survey.question2}
                      checked={response.truefalse2}
                      onChange={(value) => {
                        response.truefalse2 = value;
                        setversion(version + 1);
                        setresponses(responses);
                      }}
                    />
                  </CardContent>
                </Card>
              ))}
          </div>
        </div>
      </div>
    </div>
  );
}

/**
 * <tr key={index}>
                  
                
                    <Label htmlFor={"survey_response" + response.id}>
                      {response.displayname}
                    </Label>
             
                    <Switch
                      id={"survey_response" + response.id}
                      checked={response.truefalse1}
                      onCheckedChange={(e) => {
                        response.truefalse1 = e;
                        setversion(version + 1);
                        setresponses(responses);
                      }}
                    />
              
 */
