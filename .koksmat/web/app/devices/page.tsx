"use client";
/**
 * v0 by Vercel.
 * @see https://v0.dev/t/piUCmvToKUQ
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import Link from "next/link";
import { useService } from "@/koksmat/useservice";
import { Fragment, useContext } from "react";
import { MagicboxContext } from "@/koksmat/magicbox-context";

export interface Dashboard {
  person_id: number;
  number_of_applications: number;
  number_of_surveys: number;
  number_of_survey_responses: number;
  number_of_apps_you_own: number;
  number_of_owners: number;
  messages: any;
}
export default function Component() {
  const magicbox = useContext(MagicboxContext);
  const appDashboardService = useService<Dashboard>(
    "magic-devices.app",
    ["dashboard", magicbox.user?.email ?? "unknown"],
    "",
    600,
    "x"
  );
  return (
    <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8 lg:grid-cols-3 xl:grid-cols-3">
      <div className="grid auto-rows-max items-start gap-4 md:gap-8 lg:col-span-2">
        <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-4 lg:grid-cols-2 xl:grid-cols-4">
          <Link href="/apps/application" className="cursor-pointer">
            <Card x-chunk="dashboard-05-chunk-1">
              <CardHeader className="pb-2">
                <CardDescription>Number of applications</CardDescription>
                <CardTitle className="text-4xl">
                  {appDashboardService?.data?.number_of_applications}
                </CardTitle>
              </CardHeader>
              <CardContent>
                {/* <div className="text-xs text-muted-foreground">
            +25% from last week
          </div> */}
              </CardContent>
              <CardFooter>
                {/* <Progress value={25} aria-label="25% increase" /> */}
              </CardFooter>
            </Card>
          </Link>
          <Link href="/apps/owner" className="cursor-pointer">
            <Card x-chunk="dashboard-05-chunk-1">
              <CardHeader className="pb-2">
                <CardDescription>Number of owners</CardDescription>
                <CardTitle className="text-4xl">
                  {appDashboardService?.data?.number_of_owners}
                </CardTitle>
              </CardHeader>
              <CardContent>
                {/* <div className="text-xs text-muted-foreground">
            +25% from last week
          </div> */}
              </CardContent>
              <CardFooter>
                {/* <Progress value={25} aria-label="25% increase" /> */}
              </CardFooter>
            </Card>
          </Link>
          {appDashboardService?.data?.person_id > -1 && (
            <Fragment>
              <Link
                className="cursor-pointer"
                href={`/apps/owner/${appDashboardService?.data?.person_id}`}
              >
                <Card x-chunk="dashboard-05-chunk-1">
                  <CardHeader className="pb-2">
                    <CardDescription>Apps you own</CardDescription>
                    <CardTitle className="text-4xl">
                      {appDashboardService?.data?.number_of_apps_you_own}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    {/* <div className="text-xs text-muted-foreground">
            +25% from last week
          </div> */}
                  </CardContent>
                  <CardFooter>
                    {/* <Progress value={25} aria-label="25% increase" /> */}
                  </CardFooter>
                </Card>
              </Link>
              <Link
                href={`/apps/owner/${appDashboardService?.data?.person_id}/survey`}
                className="cursor-pointer"
              >
                <Card x-chunk="dashboard-05-chunk-1">
                  <CardHeader className="pb-2">
                    <CardDescription>Unanswered questions</CardDescription>
                    <CardTitle className="text-4xl">
                      {appDashboardService?.data?.number_of_survey_responses}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    {/* <div className="text-xs text-muted-foreground">
            +25% from last week
          </div> */}
                  </CardContent>
                  <CardFooter>
                    {/* <Progress value={25} aria-label="25% increase" /> */}
                  </CardFooter>
                </Card>
              </Link>
            </Fragment>
          )}
        </div>
      </div>
    </main>
  );
}
