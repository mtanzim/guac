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
    <div className="flex bg-slate-400 p-6 gap-8 items-center">
      <h2 className="text-xl ">Waka Dashboard</h2>
      <DatePickerWithRange date={date} setDate={setDate} />
      <Button className="mr-4" onClick={onLogout}>
        Logout
      </Button>
    </div>
  );
}
