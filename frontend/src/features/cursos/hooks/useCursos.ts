import { useCallback, useEffect, useState } from 'react';

export interface Curso { id: string; titulo: string }
interface CursosResponse { items: Curso[]; total: number }

const API_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8000';

export function useCursos(pageSize = 10) {
  const [cursos, setCursos] = useState<Curso[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPageState] = useState(1);
  const [filter, setFilterState] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const load = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const params = new URLSearchParams({ page: String(page), page_size: String(pageSize) });
      if (filter) params.set('q', filter);
      const res = await fetch(`${API_URL}/api/v1/cursos?${params}`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data: CursosResponse = await res.json();
      setCursos(data.items);
      setTotal(data.total);
    } catch (e) {
      setCursos([]);
      setTotal(0);
      setError(e instanceof Error ? e.message : 'Error desconocido');
    } finally {
      setLoading(false);
    }
  }, [page, pageSize, filter]);

  useEffect(() => { void load(); }, [load]);

  const setPage = useCallback((n: number) => setPageState(n), []);
  const setFilter = useCallback((f: string) => { setFilterState(f); setPageState(1); }, []);

  return { cursos, total, page, filter, loading, error, setPage, setFilter, refetch: load };
}