import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import { GetStaticPaths, GetStaticProps, GetStaticPropsContext } from 'next'
import { Props } from '@/lib/props'
import Meta from '@/components/meta'

const Image: React.FC<Props> = ({ text, type }) => {
  const url = process.env.baseUrl
  const router = useRouter()
  const imgURL = `${process.env.apiUrl}/api/image?Text=${text}&Type=${type}`
  useEffect(() => {
    const f = async () => router.push('/')
    f()
  }, [])
  return <Meta image={imgURL} url={url} />
}

export default Image

export const getStaticProps: GetStaticProps = async (context: GetStaticPropsContext) => {
  const { text, type } = context.params as Props
  return {
    props: { text, type },
    revalidate: 100,
  }
}

export const getStaticPaths: GetStaticPaths = async () => ({
  paths: [{ params: { text: '換気,タイム', type: 'slide' } as Props }],
  fallback: 'blocking',
})
