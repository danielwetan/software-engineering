"use client";

import ShortUrlModal from "@/components/short_url/short-url-modal";
import { getShortUrlsQuery } from "@/queries/short_url";
import { useEffect, useState } from "react";

import Link from "next/link";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export default function LinksPage() {
  // local state, no need to use redux
  const [shortUrls, setShortUrls] = useState<ShortUrl[]>([]);

  useEffect(() => {
    getShortUrls();
  }, []);

  const getShortUrls = async () => {
    try {
      const response = await getShortUrlsQuery();
      if (response.status === 200) {
        setShortUrls(response?.data);
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <>
      <div className="flex items-center justify-between">
        <h1 className="text-lg font-semibold md:text-2xl">Links</h1>
        <ShortUrlModal />
      </div>
      <div className="">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Link</TableHead>
              <TableHead className="text-right">Date</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {shortUrls?.map((item) => {
              return (
                <TableRow key={item.shortcode}>
                  <TableCell>
                    <div className="font-medium">
                      <Link
                        href={`${process.env.NEXT_PUBLIC_BASE_API}/s/${item.shortcode}`}
                        target="_blank"
                      >
                        {process.env.NEXT_PUBLIC_BASE_API}/s/{item.shortcode}
                      </Link>
                    </div>
                    <div className="hidden text-sm text-muted-foreground md:inline">
                      {item.target}
                    </div>
                  </TableCell>
                  <TableCell className="text-right">
                    {item.created_at}
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </div>
      {/* not found state */}
      {/* <div
        className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
        x-chunk="dashboard-02-chunk-1"
      >
        <div className="flex flex-col items-center gap-1 text-center">
          <h3 className="text-2xl font-bold tracking-tight">
            You have no links
          </h3>
          <p className="text-sm text-muted-foreground">
            You can start selling as soon as you add a links.
          </p>
        </div>
      </div> */}
    </>
  );
}
