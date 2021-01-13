import React, { useState, useEffect } from 'react'

import useSWR from 'swr'
import { useRouter } from 'next/router'

import fetcher from '../../lib/fetcher'
import styles from '../../styles/Home.module.css'
import Meta from '../../components/meta'

export default function Home() {
  const router = useRouter()
  const [shouldFetch, setShouldFetch] = useState(false)
  const { id } = router.query
  useEffect(() => {
    setShouldFetch(true)
  }, [id])
  const imgUrl = `${process.env.apiUrl}/api/media/${id}.png`
  const apiUrl = process.env.apiUrl
  const { error } = useSWR(
    shouldFetch ? [`${process.env.apiUrl}/api/${id}`, ''] : null,
    fetcher
  )
  return (
    <div>
      {(!error && imgUrl) && <>
        <p>Fuck</p>
        <Meta image={imgUrl} url={apiUrl} />
      </>}
      {error && router.push('/')}
    </div>
  )
}

Home.getInitialProps = ({ query }) => {
  const { id } = query
  const apiUrl = process.env.API_URL
  return {
    props: {
      id: id,
      apiUrl: apiUrl,
      imgUrl: `${apiUrl}/api/media/${id}.png`,
    },
  }
}
