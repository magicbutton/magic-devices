/**
 * v0 by Vercel.
 * @see https://v0.dev/t/piUCmvToKUQ
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { Checkbox } from "@/components/ui/checkbox";
import { Button } from "@/components/ui/button";



export default function Component() {
  return (
    <section className="w-full max-w-3xl mx-auto py-12 md:py-16 lg:py-20">
      <div className="px-4 md:px-6">
        <div className="space-y-4">
          <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl">
            Confirm Your Applications
          </h1>
          <p className="text-gray-500 dark:text-gray-400 max-w-[650px] text-lg md:text-xl">
            Select the applications you own from the list below. This will help
            us verify your account and provide the best experience.
          </p>
        </div>
        <div className="mt-8 border rounded-lg overflow-hidden">
          <table className="w-full divide-y divide-gray-200 dark:divide-gray-800">
            <thead className="bg-gray-50 dark:bg-gray-900">
              <tr>
                <th
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider dark:text-gray-400"
                  scope="col"
                >
                  Application
                </th>
                <th
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider dark:text-gray-400"
                  scope="col"
                >
                  Owned
                </th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200 dark:bg-gray-950 dark:divide-gray-800">
              <tr>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-10 w-10">
                      <img
                        alt="App Icon"
                        className="rounded-md"
                        height={40}
                        src="/placeholder.svg"
                        style={{
                          aspectRatio: "40/40",
                          objectFit: "cover",
                        }}
                        width={40}
                      />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900 dark:text-gray-50">
                        Acme CRM
                      </div>
                      <div className="text-sm text-gray-500 dark:text-gray-400">
                        Customer Relationship Management
                      </div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <Checkbox id="app-1" />
                </td>
              </tr>
              <tr>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-10 w-10">
                      <img
                        alt="App Icon"
                        className="rounded-md"
                        height={40}
                        src="/placeholder.svg"
                        style={{
                          aspectRatio: "40/40",
                          objectFit: "cover",
                        }}
                        width={40}
                      />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900 dark:text-gray-50">
                        Zeta Invoicing
                      </div>
                      <div className="text-sm text-gray-500 dark:text-gray-400">
                        Online Invoicing and Billing
                      </div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <Checkbox id="app-2" />
                </td>
              </tr>
              <tr>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-10 w-10">
                      <img
                        alt="App Icon"
                        className="rounded-md"
                        height={40}
                        src="/placeholder.svg"
                        style={{
                          aspectRatio: "40/40",
                          objectFit: "cover",
                        }}
                        width={40}
                      />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900 dark:text-gray-50">
                        Omega Project
                      </div>
                      <div className="text-sm text-gray-500 dark:text-gray-400">
                        Project Management and Collaboration
                      </div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <Checkbox id="app-3" />
                </td>
              </tr>
              <tr>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-10 w-10">
                      <img
                        alt="App Icon"
                        className="rounded-md"
                        height={40}
                        src="/placeholder.svg"
                        style={{
                          aspectRatio: "40/40",
                          objectFit: "cover",
                        }}
                        width={40}
                      />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-gray-900 dark:text-gray-50">
                        Sigma Analytics
                      </div>
                      <div className="text-sm text-gray-500 dark:text-gray-400">
                        Business Intelligence and Reporting
                      </div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <Checkbox id="app-4" />
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div className="mt-8 flex justify-end">
          <Button type="submit">Confirm Ownership</Button>
        </div>
      </div>
    </section>
  );
}
