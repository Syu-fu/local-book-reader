import React, { useState, useEffect, SyntheticEvent } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Autocomplete from "@mui/material/Autocomplete";
import Typography from "@mui/material/Typography";
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import { SubmitHandler, useForm } from "react-hook-form";
import axios from '../lib/axios'
import mc from '../lib/mc'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type BookGroup from '../types/BookGroup'
import type Book from '../types/Book'

interface BookDeleteInput {
  bookId: string
  volume: string
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

const searchBook = async (id: string, volume: string) => {
  return axios.get<Book>(
    `/book/${id}/${volume}`
  ).then((response) => {
    return response.data
  }).catch(error => {
    console.log(error);
    return undefined;
  })
}

const searchBookVolumeList = async (id: string) => {
  return axios.get<Book[]>(
    `/book/${id}`
  ).then((response) => {
    return response.data
  }).catch(error => {
    console.log(error);
    return undefined;
  })
}

const BookDeletePage = () => {
  const [searchBookGroupString, setSearchBookGroupString] = useState('');
  const { handleSubmit, register, reset } = useForm<BookDeleteInput>({ mode: 'onChange' })
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [isLoading, setIsLoading] = useState(true)
  const [options, setOptions] = useState<BookGroup[]>([]);
  const [volumeOptions, setVolumeOptions] = useState<Book[]>([]);
  const [bookId, setBookId] = useState('');
  const [volumeStr, setVolumeStr] = useState('');
  const [displayOrder, setDisplayOrder] = useState('');
  const [title, setTitle] = useState('');
  const [author, setAuthor] = useState('');
  const [publisher, setPublisher] = useState('');
  const [direction, setDirection] = useState('ltr')
  const [volumeDisabled, setVolumeDisabled] = useState(true)

  const deleteFile = (id: string, volume: string) => {
    const bucketName = "data"
    const objectsStream = mc.extensions.listObjectsV2WithMetadata(bucketName, `file/${id}/${volume}/`, true, '')
    console.log(objectsStream)
    objectsStream.on('data', async (chunk) => {
      const { name: objectName } = chunk
      await mc.removeObject(bucketName, objectName)
    });
  }


  const onSubmit: SubmitHandler<any> = (data: BookDeleteInput) => {
    setIsLoading(true)
    axios.delete<any>(
      `/book/${data.bookId}/${data.volume}`
    ).then((response) => {
      if (response.status === 200) {
        setApiMessage(`Book ${response.data.title} has been deleted.`)
        setApiError(false)
        console.log(response)
        deleteFile(data.bookId, data.volume);
      }
      else {
        setApiMessage('Delete failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Delete failed. ${error}`)
      setApiError(true)
    })
    setIsLoading(false)
  }

  const setBookVolumeList = (id: string) => {
    searchBookVolumeList(id).then(searchResults => {
      if (searchResults !== null && searchResults !== undefined) {
        setVolumeOptions([...searchResults]);
      }
    });
  }

  const selectedChange = (event: SyntheticEvent, selectedValue: BookGroup | null) => {// eslint-disable-line no-unused-vars
    if (selectedValue !== null) {
      setBookId(selectedValue.bookId)
      reset({
        bookId: selectedValue.bookId,
      })
      setBookVolumeList(selectedValue.bookId)
      setVolumeDisabled(false)
    }
  };

  const selectedChangeVolume = (event: SyntheticEvent, selectedValue: Book | null) => {// eslint-disable-line no-unused-vars
    if (selectedValue !== null) {
      setDisplayOrder(selectedValue.displayOrder.toString())
      setTitle(selectedValue.title)
      setAuthor(selectedValue.author)
      setPublisher(selectedValue.publisher)
      setDirection(selectedValue.direction)
      setIsLoading(false)
    }
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

  const debouncedBookSearchTerm = useDebounce(volumeStr, 300);
  useEffect(
    () => {
      if (debouncedBookSearchTerm !== "") {
        searchBook(bookId, volumeStr).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setVolumeOptions([])
          }
        });
      }
    },
    [debouncedBookSearchTerm]
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
            <Autocomplete
              options={volumeOptions}
              onChange={(event, value) => { return selectedChangeVolume(event, value) }}
              getOptionLabel={(option) => { return option.volume }}
              renderInput={(params) => {
                return (
                  <TextField
                    {...params}
                    id="standard-search"
                    label="Volume"
                    type="search"
                    variant="standard"
                    {...register('volume')}
                    value={volumeStr}
                    onChange={(event) => { return setVolumeStr(event.target.value) }}
                    disabled={volumeDisabled}
                  />
                )
              }}
            />
            <TextField
              id="standard-search"
              label="Display Order"
              type="search"
              variant="standard"
              value={displayOrder}
              disabled
            />
            <TextField
              id="standard-search"
              label="Title"
              type="search"
              variant="standard"
              value={title}
              disabled
            />
            <TextField
              id="standard-search"
              label="Author"
              type="search"
              variant="standard"
              value={author}
              disabled
            />
            <TextField
              id="standard-search"
              label="Publisher"
              type="search"
              variant="standard"
              value={publisher}
              disabled
            />
            <FormControl>
              <FormLabel id="direction">Direction</FormLabel>
              <RadioGroup
                row
                aria-labelledby="direction"
                name="direction"
                value={direction}
              >
                <FormControlLabel value="ltr" control={<Radio />} label="Left to Right" disabled />
                <FormControlLabel value="rtl" control={<Radio />} label="Right to Left" disabled />
              </RadioGroup>
            </FormControl>
            <LoadingButton
              size="large"
              loadingIndicator="Delete"
              variant="outlined"
              loading={isLoading}
              onClick={handleSubmit(onSubmit)}
            >
              Delete
            </LoadingButton>

            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default BookDeletePage
