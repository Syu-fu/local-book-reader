import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Search from '../pages/Search';
import BookList from '../pages/BookList';
import BookViewer from '../pages/BookViewer';
import TagAddPage from '../pages/TagAdd';
import TagEditPage from '../pages/TagEdit';
import TagDeletePage from '../pages/TagDelete';
import BookGroupAddPage from '../pages/BookGroupAdd';
import BookGroupEditPage from '../pages/BookGroupEdit';
import BookGroupDeletePage from '../pages/BookGroupDelete';
import BookAddPage from '../pages/BookAdd';
import BookEditPage from '../pages/BookEdit';
import BookDeletePage from '../pages/BookDelete';

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
          <Route path="/book/add" element={<BookAddPage />} />
          <Route path="/book/edit" element={<BookEditPage />} />
          <Route path="/book/delete" element={<BookDeletePage />} />
          <Route path="/bookgroup/add" element={<BookGroupAddPage />} />
          <Route path="/bookgroup/Edit" element={<BookGroupEditPage />} />
          <Route path="/bookgroup/delete" element={<BookGroupDeletePage />} />
        </Routes>
      </BrowserRouter >
    </div>
  )
}
export default DrawerRouter
