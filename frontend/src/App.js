import React, { useEffect, useState } from 'react';
import { useLocation } from "react-router-dom"

export default function App() {
  const [greeting, setGreeting] = useState('');
  const [colour, setColour] = useState('');

  let location = useLocation().pathname

  useEffect(() => {
    fetch(`http://localhost:8080${location}`)
      .then(response => response.json())
      .then(data => {
        setGreeting(data.Greeting)
        setColour(data.Colour)
      });
    
    document.body.style.backgroundColor = colour
  });

  return (
    <div>
      Hello {greeting}!
    </div>
  );
}