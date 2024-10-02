import { API_ENDPOINT } from "@/constants";
import { ApiCall, HTTP_METHODS } from "@/lib/network";
import { AxiosResponse } from "axios";

export const loginQuery = (
  data: LoginPayload
): Promise<AxiosResponse<LoginResponse>> => {
  return new Promise((resolve, reject) => {
    ApiCall({
      method: HTTP_METHODS.POST,
      url: API_ENDPOINT.LOGIN,
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
