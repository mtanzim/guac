import { useState } from "react";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="card bg-base-100 w-96 shadow-xl">
      <figure>
        <img
          src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg"
          alt="Shoes"
        />
      </figure>
      <div className="card-body">
        <h2 className="card-title">{count}</h2>
        <div className="card-actions justify-end">
          <button
            className="btn btn-primary"
            onClick={() => setCount((c) => c + 1)}
          >
            Count Up
          </button>
        </div>
      </div>
    </div>
  );
}

export default App;
