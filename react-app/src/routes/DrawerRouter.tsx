import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Search from '../pages/Search';
import BookList from '../pages/BookList';
import BookViewer from '../pages/BookViewer';
import TagAddPage from '../pages/TagAdd';
import TagEditPage from '../pages/TagEdit';
import TagDeletePage from '../pages/TagDelete';
import BookGroupAddPage from '../pages/BookGroupAdd';

const DrawerRouter: React.FC = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Search />} />
          <Route path="/book/:bookId/" element={<BookList />} />
          <Route path="/book/:bookId/:volume" element={<BookViewer />} />
          <Route path="/tag/add" element={<TagAddPage />} />
          <Route path="/tag/edit" element={<TagEditPage />} />
          <Route path="/tag/delete" element={<TagDeletePage />} />
          <Route path="/bookgroup/add" element={<BookGroupAddPage />} />
        </Routes>
      </BrowserRouter >
    </div>
  )
}
export default DrawerRouter
