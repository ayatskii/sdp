import { apiSlice } from './apiSlice'
import type { Media, MediaFolder } from '@/types'

export const mediaApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    // Media
    getMedia: builder.query<Media[], { folder?: number | string; type?: string; search?: string }>({
      query: (params) => ({
        url: '/media/',
        params,
      }),
      providesTags: ['Media'],
    }),
    
    getMediaItem: builder.query<Media, number>({
      query: (id) => `/media/${id}/`,
      providesTags: (_result, _error, id) => [{ type: 'Media', id }],
    }),
    
    uploadMedia: builder.mutation<Media, FormData>({
      query: (formData) => ({
        url: '/media/upload/',
        method: 'POST',
        body: formData,
      }),
      invalidatesTags: ['Media'],
    }),
    
    bulkUploadMedia: builder.mutation<Media[], FormData>({
      query: (formData) => ({
        url: '/media/bulk_upload/',
        method: 'POST',
        body: formData,
      }),
      invalidatesTags: ['Media'],
    }),
    
    updateMedia: builder.mutation<Media, { id: number; data: Partial<Media> }>({
      query: ({ id, data }) => ({
        url: `/media/${id}/`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: 'Media', id }, 'Media'],
    }),
    
    deleteMedia: builder.mutation<void, number>({
      query: (id) => ({
        url: `/media/${id}/`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Media'],
    }),
    
    bulkDeleteMedia: builder.mutation<{ message: string }, number[]>({
      query: (ids) => ({
        url: '/media/bulk_delete/',
        method: 'POST',
        body: { ids },
      }),
      invalidatesTags: ['Media'],
    }),
    
    // Media Folders
    getFolders: builder.query<MediaFolder[], { parent?: number | string }>({
      query: (params) => ({
        url: '/media-folders/',
        params,
      }),
      providesTags: ['Media'],
    }),
    
    getFolderContents: builder.query<{ subfolders: MediaFolder[]; files: Media[] }, number>({
      query: (id) => `/media-folders/${id}/contents/`,
      providesTags: ['Media'],
    }),
    
    createFolder: builder.mutation<MediaFolder, Partial<MediaFolder>>({
      query: (data) => ({
        url: '/media-folders/',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: ['Media'],
    }),
    
    updateFolder: builder.mutation<MediaFolder, { id: number; data: Partial<MediaFolder> }>({
      query: ({ id, data }) => ({
        url: `/media-folders/${id}/`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: ['Media'],
    }),
    
    deleteFolder: builder.mutation<void, number>({
      query: (id) => ({
        url: `/media-folders/${id}/`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Media'],
    }),
  }),
})

export const {
  useGetMediaQuery,
  useGetMediaItemQuery,
  useUploadMediaMutation,
  useBulkUploadMediaMutation,
  useUpdateMediaMutation,
  useDeleteMediaMutation,
  useBulkDeleteMediaMutation,
  useGetFoldersQuery,
  useGetFolderContentsQuery,
  useCreateFolderMutation,
  useUpdateFolderMutation,
  useDeleteFolderMutation,
} = mediaApi
