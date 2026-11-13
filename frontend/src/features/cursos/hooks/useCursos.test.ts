import { renderHook, act, waitFor } from '@testing-library/react';
import { describe, it, expect, vi, beforeEach } from 'vitest';
import { useCursos } from './useCursos';

const fetchMock = vi.fn();
vi.stubGlobal('fetch', fetchMock);

const okResponse = (data: unknown) => ({ ok: true, json: async () => data });

describe('useCursos Hook', () => {
  beforeEach(() => {
    fetchMock.mockReset();
  });

  it('debe manejar la carga inicial correctamente', async () => {
    const mockData = { items: [{ id: '1', titulo: 'Go' }], total: 1 };
    fetchMock.mockResolvedValueOnce(okResponse(mockData));

    const { result } = renderHook(() => useCursos());

    expect(result.current.loading).toBe(true);
    await waitFor(() => expect(result.current.loading).toBe(false));

    expect(result.current.cursos).toEqual(mockData.items);
    expect(result.current.error).toBeNull();
  });

  it('debe manejar errores de red dejando la lista vacia', async () => {
    fetchMock.mockRejectedValueOnce(new Error('Network failure'));

    const { result } = renderHook(() => useCursos());

    await waitFor(() => expect(result.current.loading).toBe(false));

    expect(result.current.error).toBe('Network failure');
    expect(result.current.cursos).toEqual([]);
  });

  it('setPage recarga con page=n y resetea a 1 si cambia el filtro', async () => {
    fetchMock.mockResolvedValue(okResponse({ items: [], total: 0 }));

    const { result } = renderHook(() => useCursos());
    await waitFor(() => expect(result.current.loading).toBe(false));

    act(() => {
      result.current.setPage(2);
    });
    await waitFor(() =>
      expect(fetchMock).toHaveBeenLastCalledWith(expect.stringContaining('page=2')),
    );

    act(() => {
      result.current.setFilter('react');
    });
    expect(result.current.page).toBe(1);
    await waitFor(() => {
      const url = fetchMock.mock.calls.at(-1)?.[0] as string;
      expect(url).toContain('page=1');
      expect(url).toContain('q=react');
    });
  });

  it('refetch vuelve a pedir los datos', async () => {
    fetchMock.mockResolvedValue(okResponse({ items: [], total: 0 }));

    const { result } = renderHook(() => useCursos());
    await waitFor(() => expect(result.current.loading).toBe(false));

    await act(async () => {
      await result.current.refetch();
    });

    expect(fetchMock).toHaveBeenCalledTimes(2);
  });
});