import { Button } from "@/components/ui/button";


export function Banner({ onLogout }: { onLogout: () => void }) {
  return (
    <div className="flex bg-slate-400 p-6 gap-8 items-center">
      <h2 className="text-xl ">Waka Dashboard</h2>
      <Button className="mr-4" onClick={onLogout}>
        Logout
      </Button>
    </div>
  );
}
