"use client";

import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { MagicboxContext } from "@/koksmat/magicbox-context";
import { useContext } from "react";

export default function Debugger() {
  const magicbox = useContext(MagicboxContext);
  return (
    <div>
      <Sheet>
        <SheetTrigger>Magic</SheetTrigger>
        <SheetContent>
          <SheetHeader>
            <SheetTitle>Debugger</SheetTitle>
          </SheetHeader>
          <SheetDescription>xx</SheetDescription>
          Service calls
          <SheetFooter>fsdf</SheetFooter>
        </SheetContent>
      </Sheet>
    </div>
  );
}
