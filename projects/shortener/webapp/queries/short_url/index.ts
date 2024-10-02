import { API_ENDPOINT } from "@/constants";
import { ApiCall, HTTP_METHODS } from "@/lib/network";
import { AxiosResponse } from "axios";

export const createShortUrlQuery = (
  data: CreateShortUrlPayload
): Promise<AxiosResponse<{}>> => {
  return new Promise((resolve, reject) => {
    ApiCall({
      method: HTTP_METHODS.POST,
      url: API_ENDPOINT.CREATE_SHORT_URL_PUBLIC,
      data,
    })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
};

export const getShortUrlsQuery = (): Promise<AxiosResponse<ShortUrl[]>> => {
  return new Promise((resolve, reject) => {
    ApiCall({
      method: HTTP_METHODS.GET,
      url: API_ENDPOINT.GET_SHORT_URL_PUBLIC,
    })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
};
