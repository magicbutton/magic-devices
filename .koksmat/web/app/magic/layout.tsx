"use client";
import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarShortcut,
  MenubarTrigger,
} from "@/components/ui/menubar";

import path from "path";
import { cwd } from "process";
import fs from "fs";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { MagicDevProvider } from "./contextprovider";

export default function Layout(props: { children: any }) {
  const { children } = props;

  return (
    <MagicDevProvider>
      <div>
        <div className="p-2">
          <Menubar>
            <MenubarMenu>
              <MenubarTrigger>MagicApp</MenubarTrigger>
              <MenubarContent>
                <MenubarItem>
                  <Link href="/" className="w-full">
                    Open
                  </Link>
                </MenubarItem>
                <MenubarItem>
                  <Link href="/magic/dev" className="w-full">
                    Develop
                  </Link>
                </MenubarItem>
              </MenubarContent>
            </MenubarMenu>

            <MenubarMenu>
              <MenubarTrigger>Services</MenubarTrigger>
              <MenubarContent>
                <MenubarItem>
                  <Link href="/magic/msaltest" className="w-full">
                    MSAL Test
                  </Link>
                </MenubarItem>
                <MenubarItem>
                  <Link href="/magic/services" className="w-full">
                    Microservices
                  </Link>
                </MenubarItem>
              </MenubarContent>
            </MenubarMenu>
          </Menubar>
        </div>
        {children}
      </div>
    </MagicDevProvider>
  );
}
