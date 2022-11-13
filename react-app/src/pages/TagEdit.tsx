import React, { useState, useEffect, SyntheticEvent } from 'react'
import TextField from "@mui/material/TextField";
import LoadingButton from '@mui/lab/LoadingButton';
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Autocomplete from "@mui/material/Autocomplete";
import Typography from "@mui/material/Typography"; import { SubmitHandler, useForm } from "react-hook-form";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import type Tag from '../types/Tag'

interface TagEditInput {
  tagId: string
  tagName: string
}

const searchCharacters = async (search: string) => {
  return axios.get<Tag[]>(
    `/tag/search/q=${search}`
  )
    .then((response) => {
      return response.data
    }).catch(error => {
      console.error(error);
      return undefined;
    });
}

const getTagByName = async (search: string) => {
  return axios.get<Tag[]>(
    `/tag/name/${search}`
  )
    .then((response) => {
      return response.data
    })
    .catch(error => {
      console.error(error);
      return undefined;
    });
}

const TagEditPage = () => {
  const [tagId, setTagId] = useState('')
  const [tagName, setTagName] = useState('');
  const [loading, setLoading] = useState(true);
  const [searchString, setSearchString] = useState('');
  const { register, handleSubmit, reset } = useForm<TagEditInput>()
  const [helperText, setHelperText] = useState('Please enter at least 1 character')
  const [isErrorTagName, setIsErrorTagName] = useState(true)
  const [apiMessage, setApiMessage] = useState('')
  const [apiError, setApiError] = useState(false)
  const [resultOptions, setResultOptions] = useState<Tag[]>([]);
  const onSubmit: SubmitHandler<TagEditInput> = (data) => {
    setLoading(true)
    axios.put<any>(
      `/tag/${data.tagId}`, {
      tagName: data.tagName
    }
    ).then((response) => {
      if (response.status === 201) {
        setApiMessage(`tagname ${response.data.tagName} has been updated.`)
        setApiError(false)
        setIsErrorTagName(true)
        setTagId('')
        setTagName('')
        reset({ tagId: "", tagName: "" })
      }
      else {
        setApiMessage('Update failed.')
        setApiError(true)
      }
    }).catch((error) => {
      setApiMessage(`Update failed. ${error}`)
      setApiError(true)
    })
  }

  const selectedChange = (event: SyntheticEvent, selectedValue: Tag | null) => {// eslint-disable-line no-unused-vars
    if (selectedValue !== null) {
      setTagId(selectedValue.tagId)
      setTagName(selectedValue.tagName)
      reset({ tagId: selectedValue.tagId, tagName: selectedValue.tagName })
    }
  };

  const debouncedSearchTerm = useDebounce(searchString, 300);
  useEffect(
    () => {
      if (debouncedSearchTerm !== "") {
        searchCharacters(debouncedSearchTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            setResultOptions([...searchResults]);
          }
        });
      }
    },
    [debouncedSearchTerm]
  );

  const debouncedGetTagTerm = useDebounce(tagName, 300);
  useEffect(
    () => {
      if (debouncedGetTagTerm !== "") {
        getTagByName(debouncedGetTagTerm).then(searchResults => {
          if (searchResults !== null && searchResults !== undefined) {
            if (searchResults.length === 0) {
              setLoading(false);
              setIsErrorTagName(false)
              setHelperText('')
            }
            else {
              setLoading(true)
              setIsErrorTagName(true)
              setHelperText('The same tag name is registered')
            }
          }
        });
      }
      else {
        setLoading(true)
        setHelperText('Please enter at least 1 character')
        setIsErrorTagName(true)
      }
    },
    [debouncedGetTagTerm]
  );

  return (
    <div>
      <ResponsiveDrawer>
        <Container maxWidth="sm" sx={{ pt: 5 }}>
          <Stack spacing={3}>
            <Autocomplete
              options={resultOptions}
              onChange={(event, value) => { return selectedChange(event, value) }}
              getOptionLabel={(option) => { return option.tagName }}
              renderInput={(params) => {
                return (
                  <TextField
                    {...params}
                    id="standard-search"
                    label="Search"
                    type="search"
                    variant="standard"
                    onChange={(event) => { return setSearchString(event.target.value) }}
                  />
                )
              }}
            />
            <TextField
              id="standard-search"
              label="Tag ID"
              type="search"
              variant="standard"
              {...register('tagId')}
              disabled
              value={tagId}
            />
            <TextField
              id="standard-search"
              label="Tag name"
              type="search"
              variant="standard"
              value={tagName}
              {...register('tagName')}
              error={isErrorTagName}
              onChange={(event) => { return setTagName(event.target.value) }}
              helperText={helperText}
            />
            <LoadingButton
              size="large"
              loading={loading}
              loadingIndicator="Edit"
              variant="outlined"
              onClick={handleSubmit(onSubmit)}
            >
              Edit
            </LoadingButton>
            <Typography color={apiError ? 'error' : ''}>{apiMessage}</Typography>
          </Stack>
        </Container>
      </ResponsiveDrawer>
    </div>
  );
}

export default TagEditPage
