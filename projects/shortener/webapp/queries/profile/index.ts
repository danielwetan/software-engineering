import { API_ENDPOINT } from "@/constants";
import { ApiCall, HTTP_METHODS } from "@/lib/network";
import { AxiosResponse } from "axios";

export const profileQuery = (): Promise<AxiosResponse<Profile>> => {
  return new Promise((resolve, reject) => {
    ApiCall({
      method: HTTP_METHODS.GET,
      url: API_ENDPOINT.PROFILE,
    })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
};
