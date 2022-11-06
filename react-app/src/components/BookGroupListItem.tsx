import React from 'react'
import ListItem from '@mui/material/ListItem'
import ListItemText from '@mui/material/ListItemText'
import Chip from "@mui/material/Chip";
import Grid from "@mui/material/Grid";
import Divider from "@mui/material/Divider";
import SellIcon from '@mui/icons-material/Sell';
import Typography from '@mui/material/Typography'
import { useNavigate } from 'react-router-dom';
import type BookGroup from '../types/BookGroup'

const BookGroupListItem: React.FC<{ bookgroup: BookGroup }> = ({bookgroup}) => {
  const navigate = useNavigate();
  const move = (nextPage: string) => {
    navigate(nextPage);
  };
  return (
    <div key={bookgroup.bookId}>
      <ListItem button key={bookgroup.bookId} onClick={() => { move(`/book/${bookgroup.bookId}`) } } >
        <Grid container >
          <Grid item xs={3}>
            <img alt={bookgroup.title} src={bookgroup.thumbnail} style={{ maxWidth: "100%" }} />
          </Grid>
          <Grid item xs={9} style={{ paddingLeft: "16px", paddingTop: "36px" }}>
            <Grid container >
              <ListItemText disableTypography primary={<Typography variant="h5" style={{ fontWeight: 'bold' }}>{bookgroup.title}</Typography>}
                secondary={<Typography style={{ marginTop: '10px' }}>{bookgroup.author}</Typography>} />
            </Grid>
            <Grid container spacing={0.5} justifyContent="flex-start">
              {bookgroup.tags.map(tag => {
                return (
                  <Grid item key={tag.tagId}>
                    <Chip icon={<SellIcon />} label={tag.tagName} variant="outlined" size="small" />
                  </Grid>
                )
              }
              )}
            </Grid>
          </Grid>
        </Grid>
      </ListItem>
      <Divider />
    </div>
  )
}
export default BookGroupListItem
