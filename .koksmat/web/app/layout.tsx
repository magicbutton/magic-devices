"use client";
import type { Metadata } from "next";
import "./globals.css";
import { MagicboxProvider } from "@/koksmat/magicbox-providers";
import { MSALWrapper } from "@/koksmat/msal/auth";
import { TailwindIndicator } from "@/components/tailwind-indicator";
import { ThemeProvider } from "@/components/theme-provider";
import Script from "next/script";
import { ServiceInspector } from "@/components/service-inspector";

export default function RootLayout2({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Script id="clarityinjection">
          {`
    (function(c,l,a,r,i,t,y){
      c[a]=c[a]||function(){(c[a].q=c[a].q||[]).push(arguments)};
      t=l.createElement(r);t.async=1;t.src="https://www.clarity.ms/tag/"+i;
      y=l.getElementsByTagName(r)[0];y.parentNode.insertBefore(t,y);
  })(window, document, "clarity", "script", "m4qodxpvid");            
            `}
        </Script>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <MagicboxProvider>
            <MSALWrapper>
              {children}
              <TailwindIndicator />
              <ServiceInspector />
            </MSALWrapper>
          </MagicboxProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
