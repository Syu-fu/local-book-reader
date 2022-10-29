import React from 'react'
import ListItem from '@mui/material/ListItem'
import ListItemText from '@mui/material/ListItemText'
import Typography from '@mui/material/Typography'
import Grid from '@mui/material/Grid'

const BookListItem: React.FC<{ src: string, title: string, author: string, volume: string }> = ({ src, title, author, volume }) => {
  return (
    <ListItem button key={volume} >
      <Grid container alignItems="center">
        <Grid item xs={3}>
          <img alt={title} src={src} style={{ maxWidth: "100%" }} />
        </Grid>
        <Grid item xs={9} style={{ paddingLeft: "16px" }} >
          <ListItemText disableTypography primary={<Typography variant="h5" style={{ fontWeight: 'bold' }}>{title}</Typography>}
            secondary={<Typography style={{ marginTop: '10px' }}>{author}</Typography>} />
        </Grid>
      </Grid>
    </ListItem>
  )
}
export default BookListItem
