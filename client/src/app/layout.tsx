import type { Metadata } from "next";
import "./globals.css";
import { TelegramProvider } from "@/app/providers/TelegramProvider";

export const metadata: Metadata = {
  title: "Quantum",
  description: "Quantum Mortgage",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body className="min-h-full bg-[#131316]">
        <TelegramProvider>
          <div className="max-w-[475px] mx-auto min-h-screen">
            {children}
          </div>
        </TelegramProvider>
      </body>
    </html>
  );
}