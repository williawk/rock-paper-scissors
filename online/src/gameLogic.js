let player1Hand = ""
let player2Hand = ""
let handTypes = ["Rock", "Paper", "Scissors"]
let winOutcomes = [["Rock", "Scissors"], ["Paper", "Rock"], ["Scissors", "Paper"]]

// function startGame(){
//   //Start game when button is pushed
// }

function game (){
  //Show three buttons: Rock, Paper, Scissors
  player1Hand = "First Button Push"
  player2Hand = "Second Button Push"
  determineWinner()
}

function determineWinner (){
  let hands = [player1Hand, player2Hand]
  if (player1Hand === player2Hand) {
    //Return draw
  } else if (winOutcomes.includes(hands)) {
    //Return Player 1 wins
  } else {
    //Return Player 2 wins
  }
}