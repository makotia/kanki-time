import React from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'

import styles from '@/styles/Share.module.css'
import { Props } from '@/lib/props'

const Share: React.FC = () => {
  const router = useRouter()
  const { text, type }: Props = router.query
  const imgURL = (text && type) ? `${process.env.apiUrl}/api/image?Text=${text}&Type=${type}` : ''
  const url = `${process.env.baseUrl}/image/${type}/${encodeURI(text)}`
  return (
    <div className={styles.container}>
      <div>
        <img className={styles.img} src={imgURL} />
        <div className={styles.share}>
          <a className={[styles.btn, styles.twBtn].join(' ')} target='_blank' href={`https://twitter.com/share?url=${encodeURI(url)}&text=オリジナルの換気タイムを作ろう！&related=0x307E&hashtags=換気タイム`}>
            Twitter にシェアする
          </a>
          <Link href="/">
            <a className={[styles.btn, styles.topBtn].join(' ')}>
              トップに戻る
            </a>
          </Link>
        </div>
      </div>
    </div>
  )
}

export default Share
