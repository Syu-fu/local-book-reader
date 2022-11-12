import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Search from '../pages/Search';
import BookList from '../pages/BookList';
import BookViewer from '../pages/BookViewer';
import TagAddPage from '../pages/TagAdd';

const DrawerRouter: React.FC = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Search />} />
          <Route path="/book/:bookId/" element={<BookList />} />
          <Route path="/book/:bookId/:volume" element={<BookViewer />} />
          <Route path="/tag/add" element={<TagAddPage />} />
        </Routes>
      </BrowserRouter >
    </div>
  )
}
export default DrawerRouter
