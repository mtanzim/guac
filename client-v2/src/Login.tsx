import { Button } from "@/components/ui/button";

import { Input } from "@/components/ui/input";
import { MouseEventHandler, useState } from "react";
import { BASE_URL } from "./utils";

export function Login({ onLogin }: { onLogin: (t: string) => void }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [errMsg, setErrMsg] = useState<string | null>(null);
  const handleSubmit: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault();
    setLoading(true);
    setErrMsg(null);

    try {
      const res = await fetch(`${BASE_URL}/api/v1/login`, {
        method: "POST",
        body: JSON.stringify({
          username,
          password,
        }),
      });
      const body = await res.json();
      if (res.ok) {
        onLogin(body?.token);
        return;
      }
      throw new Error("failed to login");
    } catch (err) {
      console.log(err);
      setErrMsg("Something went wrong");
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="w-full max-w-sm items-center">
      <Input
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        className="mt-4"
        type="username"
        placeholder="Username"
      />
      <Input
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="mt-4"
        type="password"
        placeholder="Password"
      />
      {errMsg && <p className="text-red-500">{errMsg}</p>}
      {isLoading && <p className="text-slate-400 animate-pulse">Loading...</p>}
      <Button
        className="mt-4 w-1/4 float-end"
        type="submit"
        onClick={handleSubmit}
      >
        Login
      </Button>
    </div>
  );
}
