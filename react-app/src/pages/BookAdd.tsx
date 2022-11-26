import React, { useState, useEffect, SyntheticEvent } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Autocomplete from "@mui/material/Autocomplete";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import { SubmitHandler, Controller, useForm } from "react-hook-form";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type BookGroup from '../types/BookGroup'
import { IPADDRESS, MINIO_PORT } from '../config/index'

interface BookAddInput {
  bookId: string
  volume: string
  displayOrder: number
  title: string
  author: string
  publisher: string
  direction: "ltr" | "rtl"
}

const searchBookGroup = async (search: string) => {
  return axios.get<BookGroup[]>(
    `/bookgroup/search/q=${search}`
  )
    .then((response) => {
      return response.data
    }).catch(error => {
      console.error(error);
      return undefined;
    });
}

const BookAddPage = () => {
  const [searchBookGroupString, setSearchBookGroupString] = useState('');
  const { control, handleSubmit, register, reset } = useForm<BookAddInput>({ mode: 'onChange' })
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const [options, setOptions] = useState<BookGroup[]>([]);
  const [bookId, setBookId] = useState('');
  const [direction, setDirection] = useState('ltr')

  const uploadThumbnail = (id: string, volume: string) => {
    const input = document.getElementById('thumbnail') as HTMLInputElement | null;

    if (input != null) {
      const file = input.files ? input.files[0] : null;
      if (!file) return;

      axios.put<any>(
        `http://${IPADDRESS}:${MINIO_PORT}/data/thumbnail/book/${id}/${volume}`,
        file
      )
    }
  }

  const uploadFile = (id: string, volume: string) => {
    const input = document.getElementById('file') as HTMLInputElement | null;

    if (input != null) {
      const files = input.files ? Array.from(input.files) : null;
      if (!files) return;

      files.every((file) => {
        return (
          axios.put<any>(
            `http://${IPADDRESS}:${MINIO_PORT}/data/file/${id}/${volume}/${file.name}`,
            file
          ))
      })
    }
  }

  const onSubmit: SubmitHandler<any> = (data: BookAddInput) => {
    setIsLoading(true)
    axios.post<any>(
      `/book/`, {
      bookId: data.bookId,
      volume: data.volume,
      displayOrder: Number(data.displayOrder),
      title: data.title,
      author: data.author,
      publisher: data.publisher,
      direction: data.direction,
    }
    ).then((response) => {
      if (response.status === 201) {
        setApiMessage(`Book ${response.data.title} has been created.`)
        setApiError(false)
        uploadThumbnail(response.data.bookId, response.data.volume);
        uploadFile(response.data.bookId, response.data.volume);
      }
      else {
        setApiMessage('Create failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Create failed. ${error}`)
      setApiError(true)
    })
    setIsLoading(false)
  }

  const selectedChange = (event: SyntheticEvent, selectedValue: BookGroup | null) => {// eslint-disable-line no-unused-vars
    if (selectedValue !== null) {
      setBookId(selectedValue.bookId)
      reset({
        bookId: selectedValue.bookId,
      })
    }
  };

  const handleRadio = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDirection((event.target as HTMLInputElement).value);
  };

  const debouncedBookGroupSearchTerm = useDebounce(searchBookGroupString, 300);
  useEffect(
    () => {
      if (debouncedBookGroupSearchTerm !== "") {
        searchBookGroup(debouncedBookGroupSearchTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setOptions([...searchResults]);
          }
        });
      }
    },
    [debouncedBookGroupSearchTerm]
  );

  return (
    <div>
      <ResponsiveDrawer>
        <Container maxWidth="sm" sx={{ pt: 5 }}>
          <Stack spacing={3}>
            <Autocomplete
              options={options}
              onChange={(event, value) => { return selectedChange(event, value) }}
              getOptionLabel={(option) => { return option.title }}
              renderInput={(params) => {
                return (
                  <TextField
                    {...params}
                    id="standard-search"
                    label="Search"
                    type="search"
                    variant="standard"
                    onChange={(event) => { return setSearchBookGroupString(event.target.value) }}
                  />
                )
              }}
            />
            <TextField
              id="standard-search"
              label="ID"
              type="search"
              variant="standard"
              value={bookId}
              disabled
              {...register('bookId')}
            />
            <TextField
              id="standard-search"
              label="Volume"
              type="search"
              variant="standard"
              {...register('volume')}
            />
            <TextField
              id="standard-search"
              label="Display Order"
              type="search"
              variant="standard"
              {...register('displayOrder')}
            />
            <TextField
              id="standard-search"
              label="Title"
              type="search"
              variant="standard"
              {...register('title')}
            />
            <TextField
              id="standard-search"
              label="Author"
              type="search"
              variant="standard"
              {...register('author')}
            />
            <TextField
              id="standard-search"
              label="Publisher"
              type="search"
              variant="standard"
              {...register('publisher')}
            />
            <Controller
              name="direction"
              render={() => {
                return (
                  <FormControl>
                    <FormLabel id="direction">Direction</FormLabel>
                    <RadioGroup
                      row
                      aria-labelledby="direction"
                      name="direction"
                      value={direction}
                      onChange={handleRadio}
                    >
                      <FormControlLabel value="ltr" control={<Radio />} label="Left to Right" />
                      <FormControlLabel value="rtl" control={<Radio />} label="Right to Left" />
                    </RadioGroup>
                  </FormControl>
                )
              }
              }
              control={control}
            />
            <Stack direction="row" spacing={2} justifyContent="center">
              <Typography>
                Thumbnail
              </Typography>
              <Button variant="contained" component="label">
                Upload
                <input hidden accept="image/*" type="file" id='thumbnail' />
              </Button>
            </Stack>
            <Stack direction="row" spacing={2} justifyContent="center">
              <Typography>
                File
              </Typography>
              <Button variant="contained" component="label">
                Upload
                <input hidden accept="image/*" type="file" id='file' multiple />
              </Button>
            </Stack>
            <LoadingButton
              size="large"
              loadingIndicator="Add"
              variant="outlined"
              loading={isLoading}
              onClick={handleSubmit(onSubmit)}
            >
              Add
            </LoadingButton>
            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default BookAddPage
