import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import { getSession, logout } from "@/lib/session";

const instance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_BASE_API,
  timeout: 10000,
  headers: { Accept: "application/json", "Content-Type": "application/json" },
});

instance.interceptors.request.use(
  // @ts-ignore
  // @ts-expect-error
  async (config: AxiosRequestConfig): Promise<AxiosRequestConfig> => {
    config.headers = config.headers ?? {};

    const token = await getSession();

    if (token) {
      config.headers["Authorization"] = `${token}`;
    }
    return config;
  },
  (err: AxiosError) => {
    return Promise.reject(err);
  }
);

instance.interceptors.response.use(
  (res: AxiosResponse): AxiosResponse => {
    return res;
  },
  (err: AxiosError) => {
    switch (err.response?.status) {
      case 401:
        logout();
        return Promise.reject(err);

      default:
        return Promise.reject(err);
    }
  }
);

export const HTTP_METHODS = {
  GET: "GET",
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
} as const;

type HttpMethod = keyof typeof HTTP_METHODS;

interface ApiCallConfig {
  method: HttpMethod;
  url: string;
  data?: any;
  headers?: Record<string, string>;
}

export const ApiCall = async ({
  method,
  url,
  data,
  headers = {},
}: ApiCallConfig): Promise<AxiosResponse<any>> => {
  try {
    const response = await instance({
      method,
      url,
      data,
      headers,
    });

    return response;
  } catch (error: any) {
    if (axios.isAxiosError(error)) {
      if (typeof error?.response?.data?.error === "string") {
        alert(error?.response?.data?.error);
      }
      throw error;
    } else {
      throw error;
    }
  }
};
