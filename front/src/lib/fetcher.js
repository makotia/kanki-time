export const fetcher = (url, token) =>
  fetch(url, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }).then((r) => {
    if (!r.ok) {
      throw new Error(String(r.status))
    }
    return r.json()
  })

export default fetcher
