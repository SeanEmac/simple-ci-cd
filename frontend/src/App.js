import React, { useEffect, useState } from 'react';

export default function App() {
  const [path, setPath] = useState('');
  const [colour, setColour] = useState('');
  let base = "http://localhost:8080";
  if (process.env.NODE_ENV === "production") {
    console.log('Production Mode')
    base = "";
  }

  useEffect(() => {
    fetch(`${base}/api`)
      .then(response => response.json())
      .then(data => {
        setPath(data.Path)
        setColour(data.Colour)
      });
    
    document.body.style.backgroundColor = colour
  });

  return (
    <div>
      Hello {path}!
    </div>
  );
}