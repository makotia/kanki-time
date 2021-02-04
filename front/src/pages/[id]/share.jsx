import React from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'

import styles from '../../styles/Share.module.css'

export default function Share() {
  const router = useRouter()
  const { id } = router.query
  return (
    <div className={styles.container}>
      <div>
        <img className={styles.img} src={`${process.env.apiUrl}/api/media/${id}.png`} />
        <div className={styles.share}>
          <a className={[styles.btn, styles.twBtn].join(' ')} target='_blank' href={`https://twitter.com/share?url=${encodeURI(process.env.baseUrl)}/${id}&text=オリジナルの換気タイムを作ろう！&related=0x307E`}>
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
