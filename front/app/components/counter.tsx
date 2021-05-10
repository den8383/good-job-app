import React, {useState} from 'react'

const Counter = () => {
  const [count, setCount] = useState(0);
  return (
      <div>
      <p>conunt {count}</p>
      <button onClick={() => setCount(count + 1)}>click</button>
      </div>
      )
}

export default Counter
