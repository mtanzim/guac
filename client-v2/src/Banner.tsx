import { Button } from "@/components/ui/button";
import { DatePickerWithRange } from "./DateRange";
import { DateRange } from "react-day-picker";
import { Moon, Sun } from "lucide-react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./components/ui/dropdown-menu";
import { useTheme } from "@/components/theme-provider";

export function Banner({
  onLogout,
  setDate,
  date,
}: {
  onLogout: () => void;
  setDate: (d?: DateRange) => void;
  date?: DateRange;
}) {
  return (
    <div className="flex flex-col md:flex-row bg-slate-400 p-6 gap-8 items-center">
      <div className="md:w-11/12 flex gap-4">
      <Button variant="outline" className="text-xl ">Waka Dashboard</Button>
      <ModeToggle />
      </div>
      <div className="flex flex-col md:flex-row gap-4">
        <DatePickerWithRange date={date} setDate={setDate} />
        <Button className="" onClick={onLogout}>
          Logout
        </Button>
        <a href="/v1">
          <Button variant="outline" className="w-full">
            Use v1
          </Button>
        </a>
      </div>
    </div>
  );
}

export function ModeToggle() {
  const { setTheme } = useTheme();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" size="icon">
          <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
          <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
          <span className="sr-only">Toggle theme</span>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem onClick={() => setTheme("light")}>
          Light
        </DropdownMenuItem>
        <DropdownMenuItem onClick={() => setTheme("dark")}>
          Dark
        </DropdownMenuItem>
        <DropdownMenuItem onClick={() => setTheme("system")}>
          System
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
