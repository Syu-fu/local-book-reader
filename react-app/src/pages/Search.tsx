import React, { useState, useEffect } from 'react'
import SearchIcon from "@mui/icons-material/Search";
import TextField from "@mui/material/TextField";
import axios from '../lib/axios'
import useDebounce from '../utils/useDebounce'
import ResponsiveDrawer from '../components/ResponsiveDrawer'
import BookGroupListItem from '../components/BookGroupListItem'
import type BookGroup from '../types/BookGroup'

const searchCharacters = async (search: string) => {
  if (search === "") {
    return axios.get<BookGroup[]>(
      `/bookgroup/`
    )
      .then((response) => {
        return response.data
      })
      .catch(error => {
        console.error(error);
        return undefined;
      });
  }
  return axios.get<BookGroup[]>(
    `/bookgroup/search/q=${search}`
  )
    .then((response) => {
      return response.data
    })
    .catch(error => {
      console.error(error);
      return undefined;
    });

}

const SearchBookGroupPage = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [results, setResults] = useState<BookGroup[]>([]);

  const debouncedSearchTerm = useDebounce(searchTerm, 300);

  useEffect(
    () => {
      searchCharacters(debouncedSearchTerm).then(searchResults => {
        if (searchResults !== null && searchResults !== undefined) {
          setResults(searchResults);
        }
      });
    },
    [debouncedSearchTerm]
  );
  return (
    <div>
      <ResponsiveDrawer>
        <SearchIcon />
        <TextField
          id="standard-basic"
          label="Search"
          variant="standard"
          onChange={(event) => { return setSearchTerm(event.target.value) }}
        />

        {results.map(result => {
          return (
            <BookGroupListItem bookgroup={result} />
          )
        })}
      </ResponsiveDrawer>
    </div>
  );
}

export default SearchBookGroupPage
