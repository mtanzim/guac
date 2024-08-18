import { addDays, format, endOfToday } from "date-fns";
import { Calendar as CalendarIcon } from "lucide-react";
import * as React from "react";
import { DateRange } from "react-day-picker";

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { cn } from "@/lib/utils";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export function DatePickerWithRange({
  className,
  setDate,
  date,
}: React.HTMLAttributes<HTMLDivElement> & {
  setDate: (r?: DateRange) => void;
  date?: DateRange;
}) {
  return (
    <div className={cn("grid gap-2", className)}>
      <Popover>
        <PopoverTrigger asChild>
          <Button
            id="date"
            variant={"outline"}
            className={cn(
              "w-[300px] justify-start text-left font-normal",
              !date && "text-muted-foreground"
            )}
          >
            <CalendarIcon className="mr-2 h-4 w-4" />
            {date?.from ? (
              date.to ? (
                <>
                  {format(date.from, "LLL dd, y")} -{" "}
                  {format(date.to, "LLL dd, y")}
                </>
              ) : (
                format(date.from, "LLL dd, y")
              )
            ) : (
              <span>Pick a date</span>
            )}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-auto p-1" align="start">
          <Select
            onValueChange={(value) =>
              setDate({
                to: new Date(),
                from: addDays(new Date(), parseInt(value) * -1),
              })
            }
          >
            <SelectTrigger>
              <SelectValue placeholder="Select a range" />
            </SelectTrigger>
            <SelectContent position="popper">
              <SelectItem value="7">1 week</SelectItem>
              <SelectItem value="30">1 month</SelectItem>
              <SelectItem value="90">3 months</SelectItem>
              <SelectItem value="365">1 year</SelectItem>
              <SelectItem value="1825">5 years</SelectItem>
            </SelectContent>
          </Select>
          <Calendar
            initialFocus
            mode="range"
            toDate={endOfToday()}
            defaultMonth={date?.to}
            selected={date}
            onSelect={setDate}
            numberOfMonths={2}
            className="m-2"
          />
        </PopoverContent>
      </Popover>
    </div>
  );
}

export function DatePickerWithPresets() {
  const [date, setDate] = React.useState<Date>();

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={"outline"}
          className={cn(
            "w-[280px] justify-start text-left font-normal",
            !date && "text-muted-foreground"
          )}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          {date ? format(date, "PPP") : <span>Pick a date</span>}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="flex w-auto flex-col space-y-2 p-2">
        <div className="rounded-md border">
          <Calendar mode="single" selected={date} onSelect={setDate} />
        </div>
      </PopoverContent>
    </Popover>
  );
}
