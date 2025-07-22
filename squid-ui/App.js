import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import SquidUI from './SquidUI';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/squidui" element={<SquidUI />} />
        <Route path="/" element={<Navigate to="/squidui" />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
