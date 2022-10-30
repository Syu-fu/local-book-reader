import * as React from 'react'; import List from '@mui/material/List';
import Divider from '@mui/material/Divider';
import { useParams } from 'react-router-dom'
import ResponsiveDrawer from '../components/ResponsiveDrawer';
import BookListItem from '../components/BookListItem'
import { useFetchBooks } from '../hooks/useFetchBooks'

const BookList: React.FC = () => {
  const params = useParams<{ bookId: string }>()
  console.log(params.bookId)

  const { data, error, loading } = useFetchBooks(params.bookId)

  return (
    <ResponsiveDrawer>
      <List
        sx={{
          width: '100%',
          bgcolor: 'background.paper',
        }}
      >
        <Divider variant="inset" component="li" />
        {data && data.map((book) => {return (
          <List
            sx={{
              width: '100%',
              bgcolor: 'background.paper',
            }}
          >
            <BookListItem
              src={book.thumbnail}
              title={book.title}
              author={book.author}
              volume={book.volume}
            />
            <Divider variant="inset" component="li" />
          </List>
        )})}
      </List>
      {error && <div>{error.message}</div>}
      {loading && <div>...loading</div>}
    </ResponsiveDrawer>
  )
}
export default BookList;
