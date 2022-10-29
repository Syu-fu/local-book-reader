import React from 'react';
import { BrowserRouter, Routes, Route} from "react-router-dom";
import Search from '../pages/Search';

const DrawerRouter: React.FC = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Search/>}/>
        </Routes>
      </BrowserRouter >
    </div>
  )
}
export default DrawerRouter
