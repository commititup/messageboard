import axios from 'axios';

export const apiService = {
  Get,
  Post,
  Delete,
}


function Get (url) {
  return axios.get(url)
    .then(function (response) {
        return response.data;
    })
    .catch(function (error) {
        return 'An error occured..' + error;
    })
}

function Post (url,data) {
  return axios.post(url,data)
    .then(function (response) {

        return response.data;
    })
    .catch(function (error) {
        return 'An error occured..' + error;
    })
}

function Delete (url) {
  return axios.delete(url)
    .then(function (response) {
        return response.data;
    })
    .catch(function (error) {

        return 'An error occured..' + error;
    })
}
