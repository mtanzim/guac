import { Button } from "@/components/ui/button";
import { DatePickerWithRange } from "./DateRange";
import { DateRange } from "react-day-picker";

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
      <h2 className="text-xl ">Waka Dashboard</h2>
      <div className="flex flex-col md:flex-row gap-4">
      <DatePickerWithRange date={date} setDate={setDate} />
      <Button className="" onClick={onLogout}>
        Logout
      </Button>
      </div>
    </div>
  );
}
