import { renderHook, act } from '@testing-library/react-hooks';
import useSet from './useSet';

test('useSet should accept a starting set', () => {
    // Membership is based on reference equality
    const objA = { test: 'test' };
    const objB = { test: 'test' };

    const { result } = renderHook(() => {
        const set = useSet(new Set([objA]));
        return set;
    });

    expect(result.current.has(objA)).toBeTruthy();
    expect(result.current.has(objB)).toBeFalsy();

    act(() => {
        result.current.toggle(objA);
        result.current.toggle(objB);
    });

    expect(result.current.has(objA)).toBeFalsy();
    expect(result.current.has(objB)).toBeTruthy();
});

test('useSet should correctly toggle items and report their membership', () => {
    const { result } = renderHook(() => {
        const set = useSet<string>();
        return set;
    });

    expect(result.current.has('')).toBeFalsy();
    expect(result.current.has('test')).toBeFalsy();
    expect(result.current.has('test-2')).toBeFalsy();

    act(() => {
        result.current.toggle('test');
    });

    expect(result.current.has('')).toBeFalsy();
    expect(result.current.has('test')).toBeTruthy();
    expect(result.current.has('test-2')).toBeFalsy();

    act(() => {
        result.current.toggle('test-2');
        result.current.toggle('test');
    });

    expect(result.current.has('')).toBeFalsy();
    expect(result.current.has('test')).toBeFalsy();
    expect(result.current.has('test-2')).toBeTruthy();
});
