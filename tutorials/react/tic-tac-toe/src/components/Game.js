import React from 'react';
import Board from './Board.js';
import '../index.css';
import calculateWinner from '../util.js';
class Game extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            history: [{
                squares: Array(9).fill(null),
            }],
            isXNext: true,
            stepNumber: 0,
        }
    }

    jumpTo(step) {
        this.setState({
            stepNumber: step,
            isXNext: (step % 2) === 0
        })
    }

    handleClick(i) {
        const history = this.state.history.slice(0,
            this.state.stepNumber + 1);
        const current = history[history.length - 1];
        const squares = current.squares.slice();
        if (calculateWinner(squares) || squares[i]) {
            return;
        }
        squares[i] = this.state.isXNext ? 'X' : 'O';
        this.setState({
            history: history.concat([{
                squares: squares
            }]),
            isXNext: !this.state.isXNext,
            stepNumber: history.length,
        });
    }

    render() {
        const history = this.state.history;
        const current = history[this.state.stepNumber];
        const winner = calculateWinner(current.squares);
        const moves = history.map((step, move) => {
            const desc = move ?
                'Go to move number ' + move :
                'Go to beginning';
            return (
                <li key={move}>
                    <button onClick={() =>
                        this.jumpTo(move)}>
                        {desc}
                    </button>
                </li>
            );
        })
        let status;
        if (winner) {
            status = `Winner: ${winner}`;
        } else {
            status = `Next player: ${this.state.isXNext ? 'X' : 'O'}`;
        }
        return (
            <div className="game">
                <div className="game-board">
                    <Board squares={current.squares}
                        onClick={(i) => this.handleClick(i)} />
                </div>
                <div className="game-info">
                    <div>{status}</div>
                    <ol>{moves}</ol>
                </div>
            </div>
        );
    }
}

export default Game;