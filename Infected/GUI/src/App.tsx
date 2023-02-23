import React from 'react';
import { useState } from 'react';
import reactLogo from './assets/react.svg';
import './App.css';
enum PlayerType {
  None,
  Player1,
  Player2
}
function Board({children} : {children: React.ReactNode}) {
	return (
		<div style={{
			display: 'grid',
			gridTemplateColumns: 'repeat(7, 1fr)',
			width: '600px',
			height: '600px',
		}}>
			{children}
		</div>
	);
}
type cellProps = {
  PlayerType : PlayerType,
}
function Cell(props:cellProps) {
	return (
		<div style={{
			backgroundColor: props.PlayerType === PlayerType.None ? 'white' : props.PlayerType === PlayerType.Player1 ? 'red' : 'blue',
			border: '1px solid black',
			borderRadius: '5px',
			display: 'flex',
			justifyContent: 'center',
			alignItems: 'center',
			cursor: 'pointer',
		}}>
			<div style={{
				backgroundColor: 'white',
				borderRadius: '50%',
			}}>
				{props.PlayerType === PlayerType.None ? '' : props.PlayerType === PlayerType.Player1 ? 'X' : 'O'}
			</div>
		</div>

	);
}

function App() {
	const [count, setCount] = useState(0);

	return (
		<div className="App">
			<div className='card'>
				<div className='grid'>
					<Board>
						{Array(7).fill(0).map((_, i) => {
							return Array(6).fill(0).map((_, j) => {
								return <Cell key={crypto.getRandomValues} PlayerType={PlayerType.None} />;
							});
						})
						}
					</Board>
				</div>
			</div>

		</div>
	);
}

export default App;
