import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Search from '../pages/Search';
import BookList from '../pages/BookList';

const DrawerRouter: React.FC = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Search />} />
          <Route path="/book/:bookId" element={<BookList />} />
        </Routes>
      </BrowserRouter >
    </div>
  )
}
export default DrawerRouter
