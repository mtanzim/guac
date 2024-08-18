import { Button } from "@/components/ui/button";
import { Moon, Sun } from "lucide-react";
import { DateRange } from "react-day-picker";
import { DatePickerWithRange } from "./DateRange";

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
    <div className="flex flex-col md:flex-row p-6 m-4 gap-8 items-center border-2 mb-4 rounded-md">
      <div className="md:w-11/12 flex gap-4">
        <Button variant="outline" className="text-xl ">
          Waka Dashboard
        </Button>
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
  const { setTheme, theme } = useTheme();
  const clickHandler = () => {
    if (theme == "dark") {
      setTheme("light");
      return;
    }
    setTheme("dark");
  };

  return (
    <Button onClick={clickHandler} variant="outline" size="icon">
      {theme === "dark" ? (
        <Sun className="h-[1.2rem] w-[1.2rem] transition-all" />
      ) : (
        <Moon className="absolute h-[1.2rem] w-[1.2rem] transition-all " />
      )}
    </Button>
  );
}
