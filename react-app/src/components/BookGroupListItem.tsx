import React from 'react'
import ListItem from '@mui/material/ListItem'
import ListItemText from '@mui/material/ListItemText'
import Chip from "@mui/material/Chip";
import Grid from "@mui/material/Grid";
import Divider from "@mui/material/Divider";
import SellIcon from '@mui/icons-material/Sell';
import Typography from '@mui/material/Typography'
import type BookGroup from '../types/BookGroup'

const BookGroupListItem: React.FC<{ bookgroup: BookGroup }> = ({bookgroup}) => {
  return (
    <div key={bookgroup.bookId}>
      <ListItem button>
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
                  <Grid item key={tag.TagId}>
                    <Chip icon={<SellIcon />} label={tag.TagName} variant="outlined" size="small" />
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
