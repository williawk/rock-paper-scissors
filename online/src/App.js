import './App.css';
import React, { useState } from 'react';
let choices = []


  function App() {
    const [result, setresult] = useState("Spillet har ikke startet enda!");
    const [player, setPlayer] = useState("Player1");



let gameStarts = false

let player1Hand = ""
let player2Hand = ""

let handTypes = ["Rock", "Paper", "Scissors"]
let winOutcomes = [["Rock", "Scissors"], ["Paper", "Rock"], ["Scissors", "Paper"]]

// function startGame(){
  //   //Start game when button is pushed
  // }
  
  
  
  function determineWinner (choices){
    console.log("Kommer vi hit?", choices)
    player1Hand = handTypes[choices[0]]
    player2Hand = handTypes[choices[1]]
    let hands = [player1Hand, player2Hand]
    if (player1Hand === player2Hand) {
      setresult("draw!")
      console.log("er vi her? 1")
    } else if (winOutcomes.includes(hands)) {
      setresult("Player 1 wins!")
      console.log("er vi her? 2")
    } else {
      setresult("Player 2 wins!")
      console.log("er vi her? 3")
    }
  }
  
  function setChoice(weapon){
    choices.push(weapon)
    setPlayer("Player2")
    console.log(choices)
    if (choices.length === 2){
      determineWinner(choices)
    }
  }
  



    return (
      <div className="App">
      <h1>William and Marius's 🪨 ✂️ 📃!</h1>
      <h2>⚠️BREAKING NEWS⚠️: Now online!🤩</h2>
      <h3>Choose a weapon {player}:</h3>

    <button onClick={()=>setChoice(0)}>🪨</button><button onClick={()=>setChoice(1)}>✂️</button><button onClick={()=>setChoice(2)}>📃</button>

    <p>Result:{result}</p>

    </div>
  );
}

export default App;
