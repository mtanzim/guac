import { Button } from "@/components/ui/button";
import { useState } from "react";

import { Input } from "@/components/ui/input";

export function InputWithButton() {
  return (
    <div className="w-full max-w-sm items-center">
      <Input className="mt-4" type="username" placeholder="Username" />
      <Input className="mt-4" type="password" placeholder="Password" />
      <Button className="mt-4 w-1/4 float-end" type="submit">
        Login
      </Button>
    </div>
  );
}

function App() {
  return (
    <div className="mt-64 flex flex-col justify-center items-center">
      <h2 className="text-xl">Login to Guac Dashboard</h2>
      <InputWithButton />
    </div>
  );
}

export default App;
