import { apiSlice } from './apiSlice'
import type { Site, Language, AffiliateLink, SiteFormData } from '@/types'

export const sitesApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    // Sites
    getSites: builder.query<Site[], void>({
      query: () => '/sites/',
      providesTags: ['Site'],
    }),
    getSite: builder.query<Site, number>({
      query: (id) => `/sites/${id}/`,
      providesTags: (_result, _error, id) => [{ type: 'Site', id }],
    }),
    createSite: builder.mutation<Site, SiteFormData>({
      query: (data) => ({
        url: '/sites/',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: ['Site'],
    }),
    updateSite: builder.mutation<Site, { id: number; data: Partial<SiteFormData> }>({
      query: ({ id, data }) => ({
        url: `/sites/${id}/`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: (_result, _error, { id }) => [{ type: 'Site', id }, 'Site'],
    }),
    deleteSite: builder.mutation<void, number>({
      query: (id) => ({
        url: `/sites/${id}/`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Site'],
    }),
    deploySite: builder.mutation<{ message: string; deployment_id: number }, number>({
      query: (siteId) => ({
        url: `/sites/${siteId}/deploy/`,
        method: 'POST',
      }),
      invalidatesTags: (_result, _error, siteId) => [
        { type: 'Site', id: siteId },
        'Site',
        'Deployment'
      ],
    }),
    
    // Languages
    getLanguages: builder.query<Language[], void>({
      query: () => '/languages/',
    }),
    
    // Affiliate Links
    getAffiliateLinks: builder.query<AffiliateLink[], void>({
      query: () => '/affiliate-links/',
    }),
  }),
})

export const {
  useGetSitesQuery,
  useGetSiteQuery,
  useCreateSiteMutation,
  useUpdateSiteMutation,
  useDeleteSiteMutation,
  useDeploySiteMutation,
  useGetLanguagesQuery,
  useGetAffiliateLinksQuery,
} = sitesApi
