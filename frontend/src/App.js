import React, { useEffect, useState } from 'react';
import { useLocation } from "react-router-dom"

export default function App() {
  const [path, setPath] = useState('');
  const [colour, setColour] = useState('');

  let location = useLocation().pathname;
  // let API = "http://localhost:8080";
  // if (process.env.NODE_ENV === "production") {
  //   console.log('Production Mode')
  //   API = "http://backend-service:8080";
  // }

  useEffect(() => {
    fetch(`/api${location}`)
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