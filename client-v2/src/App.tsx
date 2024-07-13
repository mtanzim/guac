import { Button } from "@/components/ui/button";

import { Input } from "@/components/ui/input";
import { MouseEventHandler, useState } from "react";

const BASE_URL = 'http://localhost:8080'


function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const handleSubmit: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault();
    

    const res = await fetch(`${BASE_URL}/api/v1/login`, {
      method: "POST",
      body: JSON.stringify({
        username,
        password,
      }),
    })
    const body = await res.json()
    alert(JSON.stringify(body));

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

function App() {
  return (
    <div className="mt-64 flex flex-col justify-center items-center">
      <h2 className="text-xl">Login to Guac Dashboard</h2>
      <Login />
    </div>
  );
}

export default App;
