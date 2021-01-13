import React, { useState, useEffect } from 'react'

import useSWR from 'swr'
import { useRouter } from 'next/router'

import fetcher from '../../lib/fetcher'
import styles from '../../styles/Home.module.css'
import Meta from '../../components/meta'

export default function Home({ id, apiUrl, imgUrl, succeed }) {
  console.log(id)
  console.log(apiUrl)
  console.log(imgUrl)
  console.log(succeed)
  return (
    <div>
      <Meta image={imgUrl} url={apiUrl} />
      {/* {!succeed && router.push('/')} */}
    </div>
  )
}

export const getStaticPaths = async () => {
  return {
    paths: [],
    fallback: true,
  }
}

export const getStaticProps = async (context) => {
  const id = context.params.id

  const { error } = useSWR(
    shouldFetch ? [`${process.env.apiUrl}/api/${id}`, ''] : null,
    fetcher
  )
  const apiUrl = process.env.API_URL

  return {
    props: {
      id: id,
      apiUrl: apiUrl,
      imgUrl: `${apiUrl}/api/media/${id}.png`,
      succeed: !error,
    },
    revalidate: 60,
  }
}
