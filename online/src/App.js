import './App.css';
import React, { useState } from 'react';
let choices = []

  function App() {
    const [result, setresult] = useState("Spillet har ikke startet enda!");
    const [player, setPlayer] = useState("Player1");

//let gameStarts = false

let player1Hand = ""
let player2Hand = ""

let handTypes = ["Rock", "Paper", "Scissors"]
let winOutcomes = [['Rock', 'Scissors'], ["Paper", "Rock"], ["Scissors", "Paper"]]

// function startGame(){
  //   //Start game when button is pushed
  // }
  
  
  
  function determineWinner (choices){
    player1Hand = handTypes[choices[0]]
    player2Hand = handTypes[choices[1]]
    let hands = JSON.stringify([player1Hand, player2Hand])

    if (player1Hand === player2Hand) {
      setresult("draw!")
      console.log("Draw")
    } else if (hands === JSON.stringify(winOutcomes[0]) || hands === JSON.stringify(winOutcomes[1]) || hands === JSON.stringify(winOutcomes[2])) {
      setresult("Player 1 wins!")
      console.log("Player 1 win")
    } else {
      setresult("Player 2 wins!")
      console.log("Player 2 win")
    }
    delay(3000).then(() => restart());
    // Need to add a time delay before showing winner
    
  }
  
  function delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
  }
  


  function restart () {
    //Show a restart button to be clicked
    //When clicked, perform the actions below
    choices = []
    setresult("Spillet har restartet!")
    setPlayer("Player 1") 
  }

  function setChoice(weapon){
    choices.push(weapon)
    setPlayer("Player 2")
    console.log(choices)
    if (choices.length === 2){
      determineWinner(choices)
    }
  }
  
    return (
      <div className="App">
      <h1>William and Marius's 🤘 📃 ✂️!</h1>
      <h2>⚠️BREAKING NEWS⚠️: Now online!🤩</h2>
      <h3>Choose a weapon {player}:</h3>

    <button onClick={()=>setChoice(0)}>🤘</button><button onClick={()=>setChoice(1)}>📃</button><button onClick={()=>setChoice(2)}>✂️</button>

    <p>Result:{result}</p>

    </div>
  );
}

export default App;
