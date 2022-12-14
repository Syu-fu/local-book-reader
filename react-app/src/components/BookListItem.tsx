import React from 'react'
import ListItem from '@mui/material/ListItem'
import ListItemText from '@mui/material/ListItemText'
import Typography from '@mui/material/Typography'
import Grid from '@mui/material/Grid'
import { useNavigate } from 'react-router-dom';
import { IPADDRESS, MINIO_PORT } from '../config/index'

const BookListItem: React.FC<{ bookId: string, title: string, author: string, volume: string }> = ({ bookId, title, author, volume }) => {
  const navigate = useNavigate();
  const move = (nextPage: string) => {
    navigate(nextPage);
  };
  return (
    <ListItem button key={volume} onClick={() => { move(`/book/${bookId}/${volume}`) }}>
      <Grid container alignItems="center">
        <Grid item xs={3}>
          <img alt={title} src={`http://${IPADDRESS}:${MINIO_PORT}/data/thumbnail/book/${bookId}/${volume}`} style={{ maxWidth: "100%" }} />
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
